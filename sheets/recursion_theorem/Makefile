# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make recursion_theorem.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

recursion_theorem.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > recursion_theorem.tex

recursion_theorem.pdf: ../../*.tex ../../trademark.pdf *.tex recursion_theorem.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode recursion_theorem.tex
	../../latexrun recursion_theorem.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: recursion_theorem.pdf
	open recursion_theorem.pdf

o:
	make open

terms: recursion_theorem.tex
	bbk sheet -mode ts recursion_theorem.tex

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
	rm -f recursion_theorem.*
	make reset
	bbk sheet -mode mk > Makefile