# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make statements.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

statements.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > statements.tex

statements.pdf: ../../*.tex ../../trademark.pdf *.tex statements.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode statements.tex
	../../latexrun statements.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: statements.pdf
	open statements.pdf

o:
	make open

terms: statements.tex
	bbk sheet -mode ts statements.tex

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
	rm -f statements.*
	make reset
	bbk sheet -mode mk > Makefile