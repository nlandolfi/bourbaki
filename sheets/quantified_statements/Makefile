# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make quantified_statements.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

quantified_statements.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > quantified_statements.tex

quantified_statements.pdf: ../../*.tex ../../trademark.pdf *.tex quantified_statements.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode quantified_statements.tex
	../../latexrun quantified_statements.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: quantified_statements.pdf
	open quantified_statements.pdf

o:
	make open

terms: quantified_statements.tex
	bbk sheet -mode ts quantified_statements.tex

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
	rm -f quantified_statements.*
	make reset
	bbk sheet -mode mk > Makefile