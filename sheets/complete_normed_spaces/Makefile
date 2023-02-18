# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make complete_normed_spaces.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

complete_normed_spaces.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > complete_normed_spaces.tex

complete_normed_spaces.pdf: ../../*.tex ../../trademark.pdf *.tex complete_normed_spaces.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode complete_normed_spaces.tex
	../../latexrun complete_normed_spaces.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: complete_normed_spaces.pdf
	open complete_normed_spaces.pdf

o:
	make open

terms: complete_normed_spaces.tex
	bbk_sheet -mode ts complete_normed_spaces.tex

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
	rm -f complete_normed_spaces.*
	make reset
	bbk_sheet -mode mk > Makefile