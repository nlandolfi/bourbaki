# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make peano_axioms.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

peano_axioms.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > peano_axioms.tex

peano_axioms.pdf: ../../*.tex ../../trademark.pdf *.tex peano_axioms.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode peano_axioms.tex
	../../latexrun peano_axioms.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: peano_axioms.pdf
	open peano_axioms.pdf

o:
	make open

terms: peano_axioms.tex
	bbk sheet -mode ts peano_axioms.tex

spell:
	aspell -c sheet.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f peano_axioms.*
	make reset
	bbk sheet -mode mk > Makefile