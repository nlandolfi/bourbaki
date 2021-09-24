# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make handwritten_digit_classifiers.pdf

handwritten_digit_classifiers.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > handwritten_digit_classifiers.tex

handwritten_digit_classifiers.pdf: ../../*.tex ../../trademark.pdf *.tex handwritten_digit_classifiers.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode handwritten_digit_classifiers.tex
	../../latexrun handwritten_digit_classifiers.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: handwritten_digit_classifiers.pdf
	open handwritten_digit_classifiers.pdf

o:
	make open

terms: handwritten_digit_classifiers.tex
	bbk_sheet -mode ts handwritten_digit_classifiers.tex

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
	rm -f handwritten_digit_classifiers.*
	make reset
	bbk_sheet -mode mk > Makefile