# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make smooth_real_functions.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

smooth_real_functions.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > smooth_real_functions.tex

smooth_real_functions.pdf: ../../*.tex ../../trademark.pdf *.tex smooth_real_functions.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode smooth_real_functions.tex
	../../latexrun smooth_real_functions.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: smooth_real_functions.pdf
	open smooth_real_functions.pdf

o:
	make open

terms: smooth_real_functions.tex
	bbk_sheet -mode ts smooth_real_functions.tex

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
	rm -f smooth_real_functions.*
	make reset
	bbk_sheet -mode mk > Makefile