# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make metric_space_functions.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

metric_space_functions.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > metric_space_functions.tex

metric_space_functions.pdf: ../../*.tex ../../trademark.pdf *.tex metric_space_functions.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode metric_space_functions.tex
	../../latexrun metric_space_functions.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: metric_space_functions.pdf
	open metric_space_functions.pdf

o:
	make open

terms: metric_space_functions.tex
	bbk_sheet -mode ts metric_space_functions.tex

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
	rm -f metric_space_functions.*
	make reset
	bbk_sheet -mode mk > Makefile