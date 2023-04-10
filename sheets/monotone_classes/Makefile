# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make monotone_classes.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

monotone_classes.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > monotone_classes.tex

monotone_classes.pdf: ../../*.tex ../../trademark.pdf *.tex monotone_classes.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode monotone_classes.tex
	../../latexrun monotone_classes.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: monotone_classes.pdf
	open monotone_classes.pdf

o:
	make open

terms: monotone_classes.tex
	bbk sheet -mode ts monotone_classes.tex

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
	rm -f monotone_classes.*
	make reset
	bbk sheet -mode mk > Makefile