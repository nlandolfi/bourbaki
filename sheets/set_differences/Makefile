# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make set_differences.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

set_differences.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > set_differences.tex

set_differences.pdf: ../../*.tex ../../trademark.pdf *.tex set_differences.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode set_differences.tex
	../../latexrun set_differences.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: set_differences.pdf
	open set_differences.pdf

o:
	make open

terms: set_differences.tex
	bbk_sheet -mode ts set_differences.tex

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
	rm -f set_differences.*
	make reset
	bbk_sheet -mode mk > Makefile