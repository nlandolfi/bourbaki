# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make function_growth_classes.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

function_growth_classes.tex: sheet.tex macros.tex
	bbk_sheet -mode c -in sheet.tex > function_growth_classes.tex

function_growth_classes.pdf: ../../*.tex ../../trademark.pdf *.tex function_growth_classes.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode function_growth_classes.tex
	../../latexrun function_growth_classes.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: function_growth_classes.pdf
	open function_growth_classes.pdf

o:
	make open

terms: function_growth_classes.tex
	bbk_sheet -mode ts function_growth_classes.tex

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
	rm -f function_growth_classes.*
	make reset
	bbk_sheet -mode mk > Makefile