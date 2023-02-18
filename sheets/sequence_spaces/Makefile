# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make sequence_spaces.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

sequence_spaces.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > sequence_spaces.tex

sequence_spaces.pdf: ../../*.tex ../../trademark.pdf *.tex sequence_spaces.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode sequence_spaces.tex
	../../latexrun sequence_spaces.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: sequence_spaces.pdf
	open sequence_spaces.pdf

o:
	make open

terms: sequence_spaces.tex
	bbk_sheet -mode ts sequence_spaces.tex

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
	rm -f sequence_spaces.*
	make reset
	bbk_sheet -mode mk > Makefile