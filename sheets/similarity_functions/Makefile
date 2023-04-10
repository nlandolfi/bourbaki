# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make similarity_functions.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

similarity_functions.tex: sheet.tex macros.tex
	bbk_sheet -mode c -in sheet.tex > similarity_functions.tex

similarity_functions.pdf: ../../*.tex ../../trademark.pdf *.tex similarity_functions.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode similarity_functions.tex
	../../latexrun similarity_functions.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: similarity_functions.pdf
	open similarity_functions.pdf

o:
	make open

terms: similarity_functions.tex
	bbk_sheet -mode ts similarity_functions.tex

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
	rm -f similarity_functions.*
	make reset
	bbk_sheet -mode mk > Makefile